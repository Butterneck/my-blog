interface Post {
  body: string,
  createdAt: Date
  description: string
  isCompleted: boolean
  title: string; 
  slug: string;
  tags: string[];
}

interface NewPost {
  title: string;
  body: string,
}