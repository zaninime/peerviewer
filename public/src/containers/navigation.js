import { connect } from 'react-redux';
import Navigation from 'components/navigation';
import { getStreamDetails } from 'reducers';

const mapStateToProps = (state, props) => {
  const { pathname } = props.location;
  let title = 'Peerviewer';
  switch (false) {
  case !/^\/watch/.test(pathname):
    {
      const details = getStreamDetails(state, props.params.streamId);
      if (!details) {
        title = 'Loading...';
        break;
      }
      let prefix;
      switch (details.mediaType) {
      case 'video':
        prefix = 'Watching';
        break;
      case 'audio':
        prefix = 'Listening';
        break;
      default:
        prefix = 'Stream';
      }
      title = `${prefix}: ${details.description}`;
    }
    break;
  }
  return { title };
};

export default connect(mapStateToProps)(Navigation);