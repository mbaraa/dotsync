<script>
	import ButtonLink from "./ButtonLink.svelte";
</script>

<section class="p-20">
	<div class="flex justify-center">
		<img
			class="w-11/12 md:w-3/4 lg:w-1/3"
			width="1000"
			height="1000"
			src="/logo.png"
			alt="Dotsync"
		/>
	</div>

	<h1 class="text-3xl pb-10 text-center">
		This is <strong>Dotsync's</strong> &quot;official&quot; documentation 📚
	</h1>

	<nav class="gap-5 text-center p-2">
		<ButtonLink href="/" title="Home" />
		<ButtonLink href="#account" title="Account Management" />
		<ButtonLink href="#add-file" title="Add a file" />
		<ButtonLink href="#add-directory" title="Add a directory" />
		<ButtonLink href="#remove-file" title="Remove a file" />
		<ButtonLink href="#remove-file" title="Remove a directory" />
		<ButtonLink href="#list-files" title="List synced files" />
		<ButtonLink href="#update-files" title="Update synced files" />
		<ButtonLink href="#download-files" title="Download synced files" />
		<ButtonLink href="#building" title="Build" />
		<ButtonLink href="#api-reference" title="API Reference" />
		<ButtonLink href="#contribution" title="Contribution" />
	</nav>

	<div class="w-full flex justify-center items-center mt-10">
		<a
			href="https://www.producthunt.com/posts/dotsync?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-dotsync"
			target="_blank"
			><img
				src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=411617&theme=light"
				alt="Dotsync - A&#0032;simple&#0032;and&#0032;concise&#0032;dotfiles&#0032;synchronizer&#0033; | Product Hunt"
				style="width: 250px; height: 54px;"
				width="250"
				height="54"
			/></a
		>
	</div>
</section>
