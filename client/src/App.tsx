import { Layout } from 'antd';

const { Sider, Header, Content, Footer } = Layout;

import './App.scss';
import { EnvironmentsTab } from './components/environment/EnvironmentsTab';

function App() {
  return (
    <Layout className="layout">
      <Header className="header">Header</Header>
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
