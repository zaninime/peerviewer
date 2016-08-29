import { combineReducers } from 'redux';
import { routerReducer as routing } from 'react-router-redux';
import streams, * as fromStreams from './streams';

export default combineReducers({routing, streams});

export const getStreamsLoaded = (state) => {
  return fromStreams.getStreamsLoaded(state.streams);
};

export const getAvailableStreams = (state) => {
  return fromStreams.getAvailableStreams(state.streams);
};

export const getStreamDetails = (state, id) => {
  return fromStreams.getStreamDetails(state.streams, id);
};
