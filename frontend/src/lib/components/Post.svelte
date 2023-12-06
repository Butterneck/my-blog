<script lang="ts">
	import { format } from 'fecha';

	export let post: Post;
</script>

<article class="flex flex-col items-start justify-between">
	<div class="w-full relative">
		<img
			src="https://images.unsplash.com/photo-1496128858413-b36217c2ce36?ixlib=rb-4.0.3&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=crop&amp;w=3603&amp;q=80"
			alt=""
			class="lg:aspect-[3/2] sm:aspect-[2/1] object-cover bg-gray-100 rounded-2xl w-full aspect-video"
		/>
		<div class="rounded-2xl ring-gray-900/10 ring-inset inset-0 absolute"></div>
	</div>
	<div class="flex items-center gap-x-4 text-xs mt-8">
		<time datetime="2020-03-16" class="text-gray-500"
			>{format(new Date(post.createdAt * 1000), 'MMM DD, YYYY')}</time
		>
		<a
			href={`/blog/${post.slug}`}
			class="relative z-10 rounded-full bg-gray-50 px-3 py-1.5 font-medium text-gray-600 hover:bg-gray-100"
			>Marketing</a
		>
	</div>
	<div class="group relative">
		<h3 class="mt-3 text-lg font-semibold leading-6 text-gray-900 group-hover:text-gray-600 line-clamp-3">
			<a href={`/blog/${post.slug}`}>
				<span class="absolute inset-0"></span>
				{post.title}
			</a>
		</h3>
		<p class="mt-5 line-clamp-3 text-sm leading-6 text-gray-600">
			<!-- {post.description} -->
			Qui natus sed id natus repellendus earum. Non quas temporibus eius. In qui necessitatibus ad. Et
			aut harum dolores eos qui reiciendis quaerat id. Sunt veniam est delectus accusamus. Suscipit eveniet
			ab voluptatibus voluptates eligendi. Deleniti fugiat dolores voluptatem cumque molestiae. Necessitatibus
			ut ut libero libero occaecati officiis.
		</p>
	</div>
	<!-- <div class="relative mt-8 flex items-center gap-x-4">
		<img
			src="https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
			alt=""
			class="h-10 w-10 rounded-full bg-gray-50"
		/>
		<div class="text-sm leading-6">
			<p class="font-semibold text-gray-900">
				<a href="#">
					<span class="absolute inset-0"></span>
					Michael Foster
				</a>
			</p>
			<p class="text-gray-600">Co-Founder / CTO</p>
		</div>
	</div> -->
</article>
