import React from 'react';

export default class Player extends React.Component {
  componentWillMount() {
    const { loaded, url, fetchStreams, goHome } = this.props;
    if (!loaded) {
      fetchStreams();
      return;
    }
    if (url === undefined) {
      goHome();
    }
  }

  render() {
    switch (this.props.mediaType) {
      case 'video':
        return (
          <video src={this.props.url} autoPlay controls preload="none">
            Your browser doesn't support the video tag!
          </video>
        );
      case 'audio':
        return (
          <audio src={this.props.url} autoPlay controls preload="none">
            Your browser doesn't support the audio tag!
          </audio>
        );
      default:
        return <div>This stream doesn't exists.</div>;
    }
  }
}

export default Player;
