<script lang="ts">
	import { invalidate } from "$app/navigation";
	import Chips from "$components/Chips/Chips.svelte";
	import viewport from "$lib/actions/viewport";
	import Carousel from "$lib/components/Carousel/Carousel.svelte";
	import Header from "$lib/components/Layouts/Header.svelte";
	import Loading from "$lib/components/Loading/Loading.svelte";
	import { homeChipContext } from "$lib/contexts";
	import type { PageData } from "./$types";
    import {APIClient} from "$lib/api";

	export let data: PageData;

	$: ({
		carousels,
		chips,
		params,
		headerThumbnail,
		continuations,
		visitorData,
		path,
	} = data);

	let loading = false;
	let hasData = false;
	$: console.log(data, carousels);
	homeChipContext.set({ params });
</script>

<svelte:head>
	{#if Array.isArray(headerThumbnail) && headerThumbnail.length !== 0 ? headerThumbnail[0].url : ""}
		<link
			rel="preload"
			as="image"
			href={headerThumbnail[0].url}
		/>
	{/if}
</svelte:head>
<Header
	title="Home"
	url={path}
	desc="Listen to the hottest tracks from your favorite artists, and discover new playlists and mixes."
/>

<div class="immersive-thumbnail">
	<div
		class="gradient"
		style="--theme: var(--base-bg);"
	/>
	{#if headerThumbnail.length !== 0}
		<picture>
			{#each headerThumbnail as thumbnail, i}
				{#if i === 0}
					<source
						media={`(max-width: ${thumbnail?.width}px)`}
						srcset={thumbnail.url}
					/>
				{:else}
					<source
						media={`(min-width: ${
							headerThumbnail[i - 1].width + 1
						}px) and (max-width: ${thumbnail?.width}px)`}
						srcset={thumbnail.url}
					/>
				{/if}
			{/each}
			<img
				src={headerThumbnail[0].url}
				width={headerThumbnail[0].width}
				height={headerThumbnail[0].height}
				decoding="async"
				class="immer-img"
				alt="large background header"
			/>
		</picture>
	{/if}
</div>
<main data-testid="home">
	<Chips
		{chips}
		on:click={() => {
			invalidate("home:load");
		}}
	/>
	{#each carousels as carousel (carousel.items)}
		<Carousel
			items={carousel.items}
			header={carousel.header}
			type="trending"
			kind={carousel.header?.type}
			isBrowseEndpoint={false}
		/>
	{/each}
	{#if Object.keys(continuations).length}
		<div
			class="viewport"
			use:viewport={{ margin: "100px" }}
			on:enterViewport={async () => {
				if (loading || hasData) return;
				loading = true;
				const response = await APIClient.fetch(
					`/api/v1/home.json?itct=${encodeURIComponent(
						continuations.clickTrackingParams,
					)}${
						params ? `&params=${encodeURIComponent(params)}` : ""
					}&ctoken=${encodeURIComponent(
						continuations.continuation,
					)}&type=next&visitorData=${visitorData}`,
				);
				const data = await response.json();
				// const {continuations, carousels} = data;
				if (data.continuations) {
					continuations = data.continuations;
					queueMicrotask(() => {
						carousels = [...carousels, ...data.carousels];
					});
					loading = false;
					return hasData;
				}
				hasData =
					data.continuations === undefined ||
					data.continuations.ctoken === undefined;
				return !loading;
			}}
		/>

		<div
			class="loading"
			style:opacity={loading ? 1 : 0}
		>
			<Loading />
		</div>
	{/if}
</main>

<style lang="scss">
	main {
		padding-bottom: 4em !important;
	}

	.viewport {
		height: 1.5em;
		position: absolute;
		bottom: 0;
		margin-top: 9rem;
		padding-block: 2.5em;
		contain: content;
		will-change: visibility;
	}

	.loading {
		transition: opacity cubic-bezier(0.95, 0.05, 0.795, 0.035) 500ms;
		opacity: 0;
		display: flex;
		position: absolute;
		// inset:0;inset
		bottom: 0;
		left: 0;
		right: 0;
		max-height: 5em;
		z-index: 100;
		margin: 0 auto;
		padding-block: 2.5rem;
		will-change: visibility;
	}

	.immersive-thumbnail {
		position: absolute;
		z-index: -1;
		width: 100%;
		max-width: 100%;
		top: 0;
		left: 0;
		right: 0;
		contain: layout style paint;
	}

	.gradient {
		position: absolute;
		inset: 0;
		width: 100%;
		height: 100%;
		background: linear-gradient(to bottom, hsl(0deg 0% 0% / 60%), var(--theme));
	}

	.immer-img {
		object-fit: cover;
		object-position: center;
		object-position: top;
		border-radius: unset !important;
		inset: 0;
		z-index: -1;
		contain: style;
		width: 100%;
		max-width: 100%;
	}
</style>
