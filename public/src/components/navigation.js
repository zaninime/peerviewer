import React from 'react';
import Radium from 'radium';
import AppBar from 'material-ui/AppBar';
import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import Divider from 'material-ui/Divider';
import IconButton from 'material-ui/IconButton';
import NavigationClose from 'material-ui/svg-icons/navigation/close';
import Home from 'material-ui/svg-icons/action/home';
import InfoOutline from 'material-ui/svg-icons/action/info-outline';
import Settings from 'material-ui/svg-icons/action/settings';
import ShowChart from 'material-ui/svg-icons/editor/show-chart';

const styles = {
  margin: '0 px',
};

const Navigation = (props) => {
  return (
    <div style={styles}>
      <AppBar
        title={props.title}
        iconClassNameRight="muidocs-icon-navigation-expand-more"
      />
      <Drawer
          docked={false}
          open={true}
          onRequestChange={null}
        >
          <AppBar
            title="Menu"
            iconElementLeft={<IconButton><NavigationClose /></IconButton>}
          />
          <MenuItem onTouchTap={null} leftIcon={<Home />}>Home</MenuItem>
          <MenuItem onTouchTap={null} leftIcon={<ShowChart />} disabled={true}>Statistics</MenuItem>
          <MenuItem onTouchTap={null} leftIcon={<Settings />} disabled={true}>Settings</MenuItem>
          <Divider/>
          <MenuItem onTouchTap={null} leftIcon={<InfoOutline />}>About</MenuItem>
        </Drawer>
      {props.children}
    </div>);
};

export default Radium(Navigation);