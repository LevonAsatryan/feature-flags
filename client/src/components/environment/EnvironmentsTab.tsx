import { Menu } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import { useEffect } from 'react';
import { fetchEnvironment } from './store/environmentsState';
import { AppDispatch } from '../../store';

import type { RootState } from '../../store';
import { Loader } from '../loader/loader';

export const EnvironmentsTab = () => {
  const dispatch = useDispatch<AppDispatch>();
  useEffect(() => {
    dispatch(fetchEnvironment());
  }, []);

  const { loading, environments } = useSelector(
    (state: RootState) => state.environments
  );

  useEffect(() => console.info(loading), [loading]);

  return (
    <>
      {loading === 'pending' && <Loader />}
      {loading === 'succeeded' && (
        <Menu theme="dark" mode="inline">
          {environments.map((env) => (
            <Menu.Item key={env.id}>
              <span className="nav-text">{env.name}</span>
            </Menu.Item>
          ))}
        </Menu>
      )}
    </>
  );
};
