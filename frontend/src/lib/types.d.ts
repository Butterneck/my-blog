interface Post {
  body: string,
  creationDate: int
  title: string; 
  slug: string;
  draft?: Draft
}

interface Draft {
  body: string,
  title: string;
}

interface NewPost {
  title: string;
  body: string,
}