// Copyright 2019 Ross Light
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import * as React from 'react';

const QUERY = gql`
query AppQuery {
  greeting
}
`;

interface QueryData {
  greeting: string;
}

export const App: React.FC<{}> = () => {
  const {data} = useQuery<QueryData>(QUERY);
  return (
    <main className="App">
      {data ?
        <p>{data.greeting}</p> :
        <p>Loading&hellip;</p>}
    </main>
  );
}
