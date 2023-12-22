<script lang="ts">
	export let href: string;
	export let title: string;
</script>

<a {href} class="w-full">
	<button
		class="rounded-lg border-2 border-[#64FFDA] p-2 hover:bg-[#64FFDA] text-[#64FFDA] hover:text-black text-2xl font-bold uppercase m-2"
		>{title}
	</button>
</a>
