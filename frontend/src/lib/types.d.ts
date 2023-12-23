interface Post {
  body: string,
  creationDate: int
  title: string; 
  slug: string;
  draft?: Draft
  assets?: string[]
}

interface Draft {
  body: string,
  title: string;
  assets?: string[]
}

interface PostAsset {
  file?: File,
  name: string,
}

interface NewPost {
  title: string;
  body: string,
  assets?: Array<Blob>;
}

interface UpdatedPost {
  slug: string;
  title: string;
  body: string;
  newAssets?: Array<Blob>;
  deletedAssets?: Array<string>;
}
