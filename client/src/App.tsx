import { Layout, Switch } from 'antd';
import { MoonOutlined, SunOutlined } from '@ant-design/icons';

const { Sider, Header, Content, Footer } = Layout;

import './App.scss';
import { EnvironmentsTab } from './components/environment/EnvironmentsTab';
import { useDispatch } from 'react-redux';
import { toggleTheme } from './store/themeState';

function App() {
  const dispatch = useDispatch();
  const switchTheme = () => {
    dispatch(toggleTheme());
  };

  return (
    <Layout className="layout">
      <Header className="header">
        <Switch
          checkedChildren={<SunOutlined />}
          unCheckedChildren={<MoonOutlined />}
          defaultChecked={false}
          onChange={switchTheme}
        />
      </Header>
      <Layout>
        <Sider className="sider">
          <EnvironmentsTab />
        </Sider>
        <Content className="content">Content</Content>
      </Layout>
      <Footer className="footer">Footer</Footer>
    </Layout>
  );
}

export default App;
