<script lang="ts">
  export let post: Post;
  import { format } from "fecha";
</script>

<div class="post-detail">
  <!-- <div class="post-detail__cover">
    <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmir-s3-cdn-cf.behance.net%2Fproject_modules%2Fmax_1200%2F0934be75683871.5c53311dabafc.png&f=1&nofb=1&ipt=8d274fb911151dead3fa383674b250eeb6084d9493869d6a6c8ce66f9777ef8a&ipo=images" alt="">
  </div> -->
  <div class="post-detail__top">
    <h1 class="post-detail__top__title">{post.title}</h1>

    <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmir-s3-cdn-cf.behance.net%2Fproject_modules%2Fmax_1200%2F0934be75683871.5c53311dabafc.png&f=1&nofb=1&ipt=8d274fb911151dead3fa383674b250eeb6084d9493869d6a6c8ce66f9777ef8a&ipo=images" alt="">

    <div class="post-detail__top__info">
      <span class="post-detail__top__info__date">
        {format(new Date(post.createdAt), "MMMM DD, YYYY")},
      </span>

      <a class="post-detail__top__info__author" href="/about">by Yourname</a>
    </div>
  </div>
  <div class="post-detail__content">
    {@html post.body}
  </div>
</div>

<style>
  .post-detail {
    padding: 1rem;
    margin-top: 1rem;
  }

  .post-detail__top {
    border-bottom: solid 1px var(--border-color);
    margin-bottom: 2rem;
  }

  .post-detail__top__title {
    padding-bottom: 2rem;
    margin: 0;
  }

  .post-detail__top__info {
    padding-bottom: 2rem;
    display: flex;
    flex-direction: row;
  }

  .post-detail__top__info__date {
    color: var(--secondary-color);
  }

  .post-detail__top__info__author {
    margin-left: 0.4rem;
  }
</style>
