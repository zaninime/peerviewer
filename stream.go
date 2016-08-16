package main

import (
	"net"
	"time"

	"github.com/zaninime/go-grapes"
	"github.com/zaninime/go-ml"
)

type rtpStream struct {
	RTP  chan []byte
	RTCP chan []byte
}

type temporalBuffer struct {
	PacketAssembly *ml.PacketAssembly
	LastUpdated    time.Time
}

type stream struct {
	assemblyBuffers   map[uint32]*temporalBuffer
	assemblyReadyChan chan []byte
	assemblyLifetime  time.Duration
	rtpStreams        []rtpStream
}

func (s *stream) ListenInbound(conn *net.UDPConn) {
	for {
		var data []byte
		_, err := conn.Read(data)
		if err != nil {
			// log
			continue // maybe break/return
		}
		packet, err := ml.ParsePacket(data)
		if err != nil {
			// log
			continue
		}
		// easy case: the packet doesn't need to be re-assembled
		if packet.ContentOffset == 0 && len(packet.Content) == int(packet.ContentTotalSize) {
			go s.handleGrapesMessage(packet.Content)
			continue
		}
		// re-assembly needed
		buf, ok := s.assemblyBuffers[packet.Sequence]
		if !ok {
			// first packet for this sequence or sequence too old
			buf = &temporalBuffer{PacketAssembly: ml.NewPacketAssembly(packet), LastUpdated: time.Now()}
			s.assemblyBuffers[packet.Sequence] = buf
		}
		buf.LastUpdated = time.Now()
		buf.PacketAssembly.Push(packet)
		if buf.PacketAssembly.Ready() {
			go s.handleGrapesMessage(buf.PacketAssembly.Buffer)
			delete(s.assemblyBuffers, packet.Sequence)
		}
	}
}

func (s *stream) handleGrapesMessage(data []byte) {
	grapesMsg, err := grapes.ParseMessage(data)
	if err != nil {
		// log
		return
	}
	switch grapesMsg.Type {
	case grapes.TypeChunk:
		s.handleChunks(grapesMsg)
	default:
		// ignore
	}
}

func (s *stream) handleChunks(msg *grapes.Message) {
	l := len(msg.Content)
	for consumed := 0; consumed < l; {
		chunk, b, err := grapes.ParseChunk(msg.Content[consumed:])
		if err != nil {
			// log
			return
		}
		s.handleRTPEnvelopes(chunk)
		consumed += int(b)
	}
}

func (s *stream) handleRTPEnvelopes(chunk *grapes.Chunk) {
	l := len(chunk.Content)
	for consumed := 0; consumed < l; {
		e, b, err := grapes.ParseRTPEnvelope(chunk.Content[consumed:])
		if err != nil {
			// log
			return
		}
		s.dispatchRTPPackets(e)
		consumed += int(b)
	}
}

func (s *stream) dispatchRTPPackets(env *grapes.RTPEnvelope) {
	if int(env.StreamID) >= len(s.rtpStreams)/2 {
		// log
		return
	}
	stream := s.rtpStreams[env.StreamID/2]
	rtp := env.StreamID%2 == 0
	if rtp {
		stream.RTP <- env.Content
	} else {
		stream.RTCP <- env.Content
	}
}

func (s *stream) CleanPartialAssemblies() {
	for seq, buf := range s.assemblyBuffers {
		if time.Now().Sub(buf.LastUpdated) > s.assemblyLifetime {
			delete(s.assemblyBuffers, seq)
		}
	}
}
