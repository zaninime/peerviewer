import { connect } from 'react-redux';
import { getStreamDetails, getStreamsLoaded } from 'reducers';
import { push } from 'react-router-redux';
import Player from 'components/Player/Player';
import { fetchStreams } from 'actions/streams';

const mapStateToProps = (state, {streamID}) => ({
  ...getStreamDetails(state, streamID),
  loaded: getStreamsLoaded(state)
});

export default connect(mapStateToProps, {goHome: () => (push('/')), fetchStreams})(Player);
