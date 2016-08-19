import React, { PropTypes as T } from 'react';
import {Link} from 'react-router';

export class Header extends React.Component {
  render() {
    const {title} = this.props;

    return (
      <div>
        <Link to="/"><h1>{title}</h1></Link>
        <section>
          Fullstack.io
        </section>
      </div>
    );
  }
}

Header.propTypes = {
  title: T.string
};

Header.defaultProps = {
  title: 'webviewer'
};

export default Header;
