import React, { PropTypes as T } from 'react';
import {Link} from 'react-router';
import AvailableStreamList from 'containers/availablestreamlist';

export class IndexPage extends React.Component {
  render() {
    return (
      <div>
        <p>
          Available streams:
        </p>
        <AvailableStreamList />
        <p>
          <Link to="about">About</Link>
        </p>
      </div>
    );
  }
}

export default IndexPage;
