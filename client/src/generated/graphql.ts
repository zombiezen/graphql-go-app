export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  readonly __typename?: 'Mutation';
  readonly mutate: Maybe<Scalars['ID']>;
};

export type MutationMutateArgs = {
  message: Scalars['String'];
};

export type Query = {
  readonly __typename?: 'Query';
  readonly greeting: Scalars['String'];
};

export type AppQueryVariables = {};

export type AppQueryData = { readonly __typename?: 'Query' } & Pick<
  Query,
  'greeting'
>;
