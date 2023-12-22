<script lang="ts">
	import CodeSnippet from "$lib/components/CodeSnippet.svelte";
	import DocsIntro from "$lib/components/DocsIntro.svelte";
	import Footer from "$lib/components/Footer.svelte";
	import Section from "$lib/components/Section.svelte";
</script>

<svelte:head>
	<title>Dotsync's Docs</title>
</svelte:head>

<main class="text-white font-Terminus">
	<DocsIntro />

	<Section id="account" title="Account Management" className="bg-[#121212]">
		<h2 class="text-xl mb-3 text-left text-white">
			You can create an account, or login into your existing account using the `<strong
				>-login</strong
			>` flag followed by a valid email of yours.
		</h2>
		<h2 class="text-xl my-3 text-white">
			For example, logging in with this email dotsync@mbaraa.com
		</h2>
		<CodeSnippet fullWidth code="dotsync -login dotsync@mbaraa.com" />

		<h2 class="text-xl mt-5 text-left text-red-600">
			Email might be delivered to spam, since it contains a <a
				class="text-[#64FFDA]"
				href="https://en.wikipedia.org/wiki/JSON_Web_Token"
				target="_blank">JWT</a
			> which some email providers consider a malicious email.
		</h2>

		<h2 class="text-xl mt-5 mb-3 text-left text-white">
			Account deletion is possible using the CLI via the `<strong>-delete-user</strong>` flag with
			no arguments, deleting your account will delete your files from the remote server,
			<strong>but it won't affect your local files</strong>.
		</h2>
		<CodeSnippet fullWidth code="dotsync -delete-user" />

		<h2 class="text-xl mt-5 text-left text-red-600">
			Account deletion will fail if you're not logged in, since there's no account to delete...
		</h2>
	</Section>

	<Section id="add-file" title="Add a file">
		<h2 class="text-xl mb-3 text-left text-white">
			To add a file to the sync list simply use the `<strong>-add</strong>` flag followed by the
			file's path,
			<br />
			Make sure that you a read/write access to the file so that it can be written on and read when synchronizing
			all files.
		</h2>
		<h2 class="text-xl mb-3 text-left text-white">
			All files are stored with their absolute path, this is to avoid any conflicts like
			<strong>~/.config/i3/config</strong>
			and <strong>~/.config/polybar/config</strong>, if both of them were added using `dotsync -add
			config` they will be stored with their absolute.
		</h2>
		<h2 class="text-xl mb-3 text-left text-white">
			<strong
				>Uploaded files have a 256KiB size limit, since it's more than enough for a dotfile, and if
				it was larger Dotsync wouldn't be free...</strong
			>
		</h2>
		<h2 class="text-xl my-3 text-white">
			For example, adding the file <strong>~/.config/i3/config</strong>
		</h2>
		<CodeSnippet fullWidth code="dotsync -add ~/.config/i3/config" />

		<h2 class="text-xl mt-3 text-left text-white">
			This will add all sub files and directories inside of <strong>~/.config/i3</strong>, so
			careful with the files' sizes!
		</h2>

		<h2 class="text-xl mt-5 text-left text-red-600">
			Adding a file will fail if you're not logged in!
		</h2>
	</Section>

	<Section id="add-directory" title="Add a directory" className="bg-[#121212]">
		<h2 class="text-xl mb-3 text-left text-white">
			Adding a directory to the sync list is similar to adding a file using the `<strong
				>-add</strong
			>` flag followed by the directory's path.
		</h2>
		<h2 class="text-xl mb-3 text-left text-white">
			<strong>The size limit of a directory is 50MiB with each file not larger than 256KiB.</strong>
			<br />
		</h2>
		<h2 class="text-xl my-3 text-white">
			For example, adding the directory <strong>~/.config/i3</strong>
		</h2>
		<CodeSnippet fullWidth code="dotsync -add ~/.config/i3/" />

		<h2 class="text-xl mt-5 text-left text-red-600">
			Adding a directory will fail if you're not logged in!
		</h2>
	</Section>

	<Section id="remove-file" title="Remove a file">
		<h2 class="text-xl mb-3 text-left text-white">
			To remove a file from the sync list simply use the `<strong>-remove</strong>` flag followed by
			the file's path (same as in the sync list),
			<br />
			Make sure that the file exists in the sync list.
		</h2>
		<h2 class="text-xl my-3 text-white">
			For example, removing the file <strong>~/.config/i3/config</strong>
		</h2>
		<CodeSnippet fullWidth code="dotsync -remove ~/.config/i3/config" />

		<h2 class="text-xl mt-5 text-left text-red-600">
			Deleting a file will fail if you're not logged in!
		</h2>
	</Section>

	<Section id="remove-directory" title="Remove a Directory" className="bg-[#121212]">
		<h2 class="text-xl mb-3 text-left text-white">
			Same as files, but with a directory name, where the server will do its magic!
			<h2 />
		</h2></Section
	>

	<Section id="list-files" title="List Synced Files">
		<h2 class="text-xl mb-3 text-left text-white">
			To list current synced files use the `<strong>-list</strong>` flag with no arguments.
		</h2>
		<CodeSnippet fullWidth code="dotsync -list" />

		<h2 class="text-xl mt-3 text-left text-white">
			This will list all synced files in a chronological order, with their absolute path, so that
			you know what is synced.
		</h2>

		<h2 class="text-xl mt-5 text-left text-red-600">
			Listing synced files will fail if you're not logged in!
		</h2>
	</Section>

	<Section id="update-files" title="Upload Synced Files" className="bg-[#121212]">
		<h2 class="text-xl mb-3 text-left text-white">
			To synchronize files from your computer to the remote server use the `<strong>-upload</strong
			>` flag with no arguments.
		</h2>
		<CodeSnippet fullWidth code="dotsync -upload" />

		<h2 class="text-xl mt-3 text-left text-white">
			Uploading changed files follows the same rules that apply for uploading a single file,
			<strong>that is each updated file should not be larger than 256KiB.</strong>
		</h2>

		<h2 class="text-xl mt-5 text-left text-red-600">
			Uploading files will fail if you're not logged in!
		</h2>
	</Section>

	<Section id="download-files" title="Download Synced Files">
		<h2 class="text-xl mb-3 text-left text-white">
			To download currently synced files from the remote server to your computer use the `<strong
				>-download</strong
			>` flag with no arguments.
		</h2>
		<CodeSnippet fullWidth code="dotsync -download" />

		<h2 class="text-xl mt-3 text-left text-white">
			When downloading synced files, if a file doesn't exist Dotsync will create a file with the
			remote server's version, so this can be very handy with fresh installs, just saying 👀
		</h2>

		<h2 class="text-xl mt-5 text-left text-red-600">
			Downloading synced files will fail if you're not logged in!
		</h2>
	</Section>

	<Section
		id="building"
		title="Building Guide"
		subTitle="Building the CLI, you need to have Go installed!"
		className="bg-[#121212]"
	>
		<h2 class="text-xl mb-3 text-left text-white">
			A handy <a class="text-[#64FFDA]" href="https://www.gnu.org/software/make" target="_blank"
				>Makefile</a
			> is included to do a proper build from source.
		</h2>
		<h2 class="text-xl my-3 text-white">
			Just compile the project, this will take no time, given the fact that the Go compiler is
			blaznigly fast.
		</h2>
		<CodeSnippet fullWidth code="make" />

		<h2 class="text-xl my-3 text-white">Installing the compiled binary.</h2>
		<CodeSnippet fullWidth code="sudo make install" />

		<h2 class="text-xl my-3 text-white">Compiling and installing the latest release from the remote repository.</h2>
		<CodeSnippet fullWidth code="sudo make install_remote" />
	</Section>

	<Section id="api-reference" title="API Reference">
		<h2 class="text-xl mb-3 text-left text-white">
			For now just open the <a
				class="text-[#64FFDA]"
				href="https://github.com/mbaraa/dotsync_server"
				target="_blank">server</a
			>'s repo and check the controllers...
		</h2>
	</Section>

	<Section id="contribution" title="Contribution" className="bg-[#121212]">
		<h2 class="text-xl mb-3 text-left text-white">
			Do it if you want, start with this section please!
		</h2>
	</Section>

	<Footer />
</main>
