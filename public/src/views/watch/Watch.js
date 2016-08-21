import React, {PropTypes as T} from 'react';
import Player from 'containers/Player/Player';

const Watch = ({params}) => {
  return (
    <div>
      Hello world! Stream: {params.streamID}

      <div>
        <Player streamID={params.streamID}/>
      </div>
    </div>
  );
};

export default Watch;
