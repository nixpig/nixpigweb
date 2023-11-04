export type Post = {
  id: string;
  title: string;
  subtitle: string;
  body: string;
  slug: string;
  status: number;
  created_at: Date;
  published_at: Date;
  updated_at: Date;
  user_id: number;
  category_id: number;
};

export type NewPost = {
  user_id: number;
  title: string;
  subtitle: string;
  body: string;
  status: number;
  category_id: number;
};
