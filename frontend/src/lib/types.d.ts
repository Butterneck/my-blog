interface Post {
  body: string,
  creationDate: int
  title: string; 
  slug: string;
  draft?: Draft
  attachments?: PostAttachment[]
}

interface Draft {
  body: string,
  title: string;
  attachments?: PostAttachment[]
}

interface NewPost {
  title: string;
  body: string,
}

interface PostAttachment {
  name: string;
  url?: string;
}