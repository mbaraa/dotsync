<script lang="ts">
	import Footer from "$lib/components/Footer.svelte";
	import Intro from "$lib/components/Intro.svelte";
	import Section from "$lib/components/Section.svelte";
	import { onMount } from "svelte";
	import CodeSnippet from "$lib/components/CodeSnippet.svelte";

	let releaseTag = "latest";

	async function getLatestRelease() {
		releaseTag = await fetch("/api/latest-release")
			.then((resp) => resp.json())
			.then((resp) => resp["tag"]);
	}

	onMount(async () => {
		await getLatestRelease();
	});
</script>

<svelte:head>
	<title>Dotsync</title>
</svelte:head>

<main class="text-white font-Terminus">
	<Intro />

	<Section id="features" title="Features✨" className="bg-[#121212]">
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 place-items-center gap-10">
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Efficient</h2>
				<p class="text-[#dedede] text-lg">
					The shining point in Dotsync. synchronizing dotfiles in a seamless way, instead of copying
					dotfile back and forth to/from a Git repository.
				</p>
			</div>
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Lightweight</h2>
				<p class="text-[#dedede] text-lg">
					The standalone CLI is only <strong>5MiB</strong>, and takes &LessTilde;
					<strong>10MiB</strong> of memory while running!
				</p>
			</div>
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Blazingly Fast</h2>
				<p class="text-[#dedede] text-lg">
					Well that's a bit stretching it, but since it’s really simple and straight to the point,
					so performance is just snappy!
				</p>
			</div>
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Cool Stack</h2>
				<p class="text-[#dedede] text-lg">
					This is my favorite thing, Dotsync has a very fancy and diverse stack, where The <a
						class="text-[#64FFDA]"
						href="https://github.com/mbaraa/dotsync"
						target="_blank">CLI</a
					>
					is built with <a class="text-[#64FFDA]" href="https://golang.org" target="_blank">Go</a>,
					the
					<a class="text-[#64FFDA]" href="https://github.com/mbaraa/dotsync_server" target="_blank"
						>server</a
					>
					with <a class="text-[#64FFDA]" href="https://elixir-lang.org" target="_blank">Elixir</a>
					&amp;
					<a class="text-[#64FFDA]" href="https://phoenixframework.org" target="_blank">Phoenix</a>,
					and the
					<a class="text-[#64FFDA]" href="https://github.com/mbaraa/dotsync_website" target="_blank"
						>website</a
					>
					with <a class="text-[#64FFDA]" href="https://kit.svelte.dev" target="_blank">SvelteKit</a>
					&amp;
					<a class="text-[#64FFDA]" href="https://tailwindcss.com" target="_blank">Tailwind CSS</a>.
				</p>
			</div>
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Open-source</h2>
				Dotsync is an open-source project licensed under
				<a
					class="text-[#64FFDA]"
					href="https://www.gnu.org/licenses/gpl-3.0.en.html"
					target="_blank">GPL-3.0</a
				>, you can star it, fork it, open an issue, or make a pull request at any repository you
				desire &#10100;<a
					class="text-[#64FFDA]"
					href="https://github.com/mbaraa/dotsync"
					target="_blank">CLI</a
				>&comma;
				<a class="text-[#64FFDA]" href="https://github.com/mbaraa/dotsync_server" target="_blank"
					>server</a
				>&comma;
				<a class="text-[#64FFDA]" href="https://github.com/mbaraa/dotsync_website" target="_blank"
					>website</a
				>&#x2775;
				<p class="text-[#dedede] text-lg" />
			</div>
			<div class="w-80">
				<h2 class="text-xl font-bold text-white text-center mb-3">Free</h2>
				<p class="text-[#dedede] text-lg">
					Dotsync is completely free (as in both freedom and charge), but I wouldn't mind if you got
					me a <a class="text-[#64FFDA]" href="https://www.buymeacoffee.com/mbaraa" target="_blank"
						>coffee</a
					>.
				</p>
			</div>
		</div>
	</Section>

	<Section
		id="installation"
		title="Installation⬇️"
	>
	    <h2 class="text-2xl pb-10 text-center text-[#64FFDA] font-bold">Use the Go installer for a quick installaion.</h2>
		<CodeSnippet code={`go install github.com/mbaraa/dotsync@${releaseTag}`} />

	    <h2 class="text-2xl py-10 text-center text-[#64FFDA] font-bold">Add Go's bin path to your path.</h2>
		<CodeSnippet code={'echo "export PATH=$HOME/go/bin:$PATH" >> ~/.`basename $SHELL`rc'} />
	</Section>

	<Section
		id="quick-start"
		title="Quick Start🚀"
		subTitle="A quick guide to backup a file."
		className="bg-[#121212]"
	>
		<h2 class="text-xl mt-5 mb-3 text-white font-bold">
			1. Login using an email, and follow the steps
		</h2>
		<CodeSnippet fullWidth code="dotsync -login someone@example.com" />

		<h2 class="text-xl mt-5 mb-3 text-white font-bold">
			2. Add a file to your sync list, for example <strong>~/.bashrc</strong>
		</h2>
		<CodeSnippet fullWidth code="dotsync -add ~/.bashrc" />

		<h2 class="text-xl mt-5 mb-3 text-white font-bold">
			3. Login into another computer
			<br />
			4. Sync the files on the other computer
		</h2>
		<CodeSnippet fullWidth code="dotsync -download" />

		<h2 class="text-xl mt-5 mb-3 text-white font-bold">
			5. Update your files after a local change
		</h2>
		<CodeSnippet fullWidth code="dotsync -upload" />

		<h2 class="text-2xl mt-10 mb-3 text-center text-white font-bold">
			For a more detailed usage visit the <a class="text-[#64FFDA]" href="/docs">Docs</a>
		</h2>
	</Section>

	<Section id="privacy" title="Privacy🔒" subTitle="How private and secure is Dotsync?">
		<h2 class="text-xl my-1 text-left text-white font-bold w-full">
			1. Dotsync doesn't include any kind of data telemetry trackers, your data are encrypted inside
			of a docker container, so chances of a leakage are near zero!
		</h2>
		<h2 class="text-xl my-3 text-left text-white font-bold w-full">
			2. I herby promise to never look at any of the uploaded data, I don't know if it means
			anything, but trust me, I ain't got time to browse containers and decrypt files manually.
		</h2>
		<h2 class="text-xl my-3 text-left text-white font-bold w-full">
			3. Finally, as you can see Dotsync is open source, and you can check how the data flows across
			the application.
		</h2>
	</Section>

	<Footer />
</main>
