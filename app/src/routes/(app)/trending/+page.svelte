<script lang="ts">
	import { browser } from "$app/environment";
	import Carousel from "$components/Carousel/Carousel.svelte";
	import Header from "$lib/components/Layouts/Header.svelte";


	export let data;
	const { carousels, page: path } = data;
	$: browser && console.log(data);
</script>

<Header
	title="Trending"
	url={path}
	desc="The latest trending songs and releases"
/>
<main data-testid="trending">
	{#each carousels as carousel (carousel)}
		{#if carousel.items}
			<Carousel
				isBrowseEndpoint={false}
				header={carousel.header}
				items={carousel.items}
				type="trending"
				kind="isPlaylist"
			/>
		{:else if carousel.categories}
			<div class="breakout">
				<div class="header">
					<span class="h2">{carousel.header.title}</span>
					<a
						class="link"
						href="/explore"><small>See All</small></a
					>
				</div>
				<div class="box">
					<div class="scroll">
						{#each carousel.categories as item}
							<a
								style="--color: {item?.color}"
								class="item-box"
								href="/explore/{item?.endpoint?.params}">{item?.text}</a
							>
						{/each}
					</div>
				</div>
			</div>
		{/if}
	{/each}
</main>

<style lang="scss">
	a small {
		$color: rgb(175 175 175);

		font-size: 0.95rem;
		font-weight: 700;
		font-variant-caps: all-petite-caps;
		letter-spacing: 0.05rem;
		transition: ease-in color 75ms;
		color: $color;

		&:hover {
			color: lighten($color, 30%);
			text-decoration: underline 0.001rem solid;
			text-underline-offset: 0.001rem;
		}
	}

	.breakout {
		border-radius: 0.8rem;
		-webkit-overflow-scrolling: touch;
		position: relative;
		margin-bottom: 2rem;

		@media screen and (min-width: 960px) {
			margin-bottom: 3rem;
		}
	}

	.box {
		display: flex;
		width: 100%;
		overflow-x: auto;
		padding: 0.8rem;
		contain: content;
		flex-direction: column;
	}

	.scroll {
		display: flex;
		flex-flow: column wrap;
		gap: 0.8rem;
		justify-content: space-around;
		//
		max-height: calc(100vh - 1px - calc(100vh - 23em));
	}

	.item-box {
		cursor: pointer;
		background: #2727298a;
		display: flex;
		justify-content: flex-start;
		flex-flow: row nowrap;
		text-overflow: clip;
		width: clamp(12em, 13em, 15em);
		contain: content;
		border-radius: 0.3em;
		font-family: "Commissioner Variable", sans-serif;
		border-left: 0.5rem solid var(--color, red);
		align-items: center;
		height: 3.25em;
		padding: 0 0 0 0.8rem;
	}
</style>
