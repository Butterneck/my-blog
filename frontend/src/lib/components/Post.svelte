<script lang="ts">
  export let post: Post;
  import { format } from "fecha";
</script>

<div class="post-item">
  <!-- <p class="post-item__date">{format(new Date(post.createdAt * 1000), "MMMM DD, YYYY")}</p> -->
  <img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmir-s3-cdn-cf.behance.net%2Fproject_modules%2Fmax_1200%2F0934be75683871.5c53311dabafc.png&f=1&nofb=1&ipt=8d274fb911151dead3fa383674b250eeb6084d9493869d6a6c8ce66f9777ef8a&ipo=images" alt="">
  <h2 class="post-item__title">
    <a class="post-item__title__link" href={`/blog/${post.slug}`}
      >{post.title}</a
    >
  </h2>
  <!-- <p class="post-item__description">{post.description}</p> -->
</div>

<style>
  .post-item {
    border-bottom: solid 1px var(--border-color);
    padding: 1rem 0;
    line-height: 1.6;
    text-align: center;
  }
  .post-item__date {
    color: var(--secondary-color);
  }

  .post-item__title {
    margin: 0.5rem 0;
  }

  .post-item__title__link {
    color: var(--main-color);
    font-size: 1.5rem;
    text-decoration: none;
  }
</style>
