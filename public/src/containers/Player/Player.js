import { connect } from 'react-redux';
import { getStreamDetails } from 'reducers';
import { push } from 'react-router-redux';
import Player from 'components/Player/Player';

const mapStateToProps = (state, {streamID}) => ({
  ...getStreamDetails(state, streamID)
});

export default connect(mapStateToProps, {goHome: () => (push('/'))})(Player);
