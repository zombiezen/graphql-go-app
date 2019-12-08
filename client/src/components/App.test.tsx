// Copyright 2019 Ross Light
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { MockedProvider, MockedResponse } from '@apollo/react-testing';
import { mount } from 'enzyme';
import * as React from 'react';
import { act } from 'react-dom/test-utils';
import waitForExpect from 'wait-for-expect';

import { App, QUERY } from './App';
import { MemoryRouter } from 'react-router';

describe('App', () => {
  const mocks: MockedResponse[] = [
    {
      request: {
        query: QUERY,
      },
      result: {
        data: {
          greeting: 'Hello, World!',
        },
      },
    },
  ];
  const Wrapper: React.FC<{children?: any}> = ({children}) => (
    <MockedProvider mocks={mocks} addTypename={false}>
      <MemoryRouter initialEntries={[ { pathname: '/', key: 'xyzzy' } ]}>
        {children}
      </MemoryRouter>
    </MockedProvider>
  );

  it('renders a loading message', () => {
    const component = mount(<App />, {
      wrappingComponent: Wrapper,
    });
    expect(component).toMatchSnapshot();
  });
  it('renders the greeting when loaded', async () => {
    const component = mount(<App />, {
      wrappingComponent: Wrapper,
    });
    await act(async () => {
      await waitForExpect(() => {
        component.update();
        expect(component.text()).toEqual(
          expect.stringContaining('Hello, World!'));
      });
    });
    expect(component).toMatchSnapshot();
  });
});
