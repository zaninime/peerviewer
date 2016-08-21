import React, {PropTypes as T} from 'react';
import { Link } from 'react-router';

export default class StreamList extends React.Component {
  componentDidMount() {
    this.props.fetchStreams();
  }

  render() {
    const streams = this.props.streams.map((stream) => {
      return (
        <li key={stream.id}>
          <Link to={`/watch/${stream.id}`}>
            {stream.description} ({stream.mediaType})
          </Link>
        </li>
      );
    });
    return (
      <ul>{streams}</ul>
    );
  }
}
