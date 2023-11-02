export type User = {
  id: number;
  username: string;
  is_admin: boolean;
  email: string;
  password?: string;
  registered_at: Date;
  last_login: Date;
  role: string;
  profile: string;
};

export type NewUser = {
  username: string;
  email: string;
  password: string;
};
