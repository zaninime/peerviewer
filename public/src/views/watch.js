import React, {PropTypes as T} from 'react';
import Player from 'containers/player';

const Watch = ({params}) => {
  return (
    <div>
      Hello world! Stream: {params.streamId}

      <div>
        <Player streamId={params.streamId}/>
      </div>
    </div>
  );
};

export default Watch;
