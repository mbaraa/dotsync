<script lang="ts">
	export let title: string;
	export let subTitle: string = "";
	export let id: string;
	export let className: string = "";
</script>

<section {id} class={"h-auto p-10 " + className}>
	<h1 class="text-4xl pb-10 text-center text-[#64FFDA] font-bold">{title}</h1>
    {#if subTitle.length > 0}
	    <h2 class="text-2xl pb-10 text-center text-[#64FFDA] font-bold">{subTitle}</h2>
    {/if}

	<slot />
</section>
