import { Menu } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import { useEffect } from 'react';
import { fetchEnvironment } from '../../store/environmentsState';
import { AppDispatch } from '../../store';

import type { RootState } from '../../store';
import { Loader } from '../loader/loader';

import './EnvironmentTab.scss';
import { fetchFeatureFlags } from '../../store/featureFlagsState';

export const EnvironmentsTab = () => {
  const { theme } = useSelector((state: RootState) => state.theme);
  const dispatch = useDispatch<AppDispatch>();
  useEffect(() => {
    dispatch(fetchEnvironment());
  }, []);

  const { loading, environments } = useSelector(
    (state: RootState) => state.environments
  );

  const getFeatureFlags = (envId: string) => {
    dispatch(fetchFeatureFlags(envId));
  };

  return (
    <>
      {loading === 'pending' && <Loader />}
      {loading === 'succeeded' && (
        <Menu mode="inline" theme={theme} className="menu">
          {environments.map((env) => (
            <Menu.Item key={env.id}>
              <span
                className="nav-text"
                onClick={() => getFeatureFlags(env.id)}
              >
                {env.name}
              </span>
            </Menu.Item>
          ))}
        </Menu>
      )}
    </>
  );
};
