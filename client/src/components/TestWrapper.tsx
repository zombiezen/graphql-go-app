import { MockedProvider, MockedResponse } from '@apollo/react-testing';
import React from 'react';
import { MemoryRouter } from 'react-router-dom';

export interface TestWrapperProps {
  mocks?: MockedResponse[];
}

export const TestWrapper: React.FC<TestWrapperProps> = ({
  mocks,
  children,
}) => (
  <MockedProvider mocks={mocks} addTypename={false}>
    <MemoryRouter initialEntries={[{ pathname: '/', key: 'xyzzy' }]}>
      {children}
    </MemoryRouter>
  </MockedProvider>
);
