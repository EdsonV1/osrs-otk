<script lang="ts">
    import { onMount } from 'svelte';
    
    // Performance-optimized icon component
    export let src: string;
    export let alt: string;
    export let size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' = 'md';
    export let loading: 'lazy' | 'eager' = 'lazy';
    export let preload: boolean = false;
    export let fallback: string = '';
    export let classes: string = '';
    
    // Size mappings for consistent performance
    const sizeMap = {
        xs: 'w-4 h-4',
        sm: 'w-6 h-6', 
        md: 'w-8 h-8',
        lg: 'w-10 h-10',
        xl: 'w-12 h-12'
    };
    
    // Image cache for performance
    let imgElement: HTMLImageElement;
    let isLoaded = false;
    let hasError = false;
    
    // Preload critical icons
    onMount(() => {
        if (preload) {
            const preloadLink = document.createElement('link');
            preloadLink.rel = 'preload';
            preloadLink.href = src;
            preloadLink.as = 'image';
            document.head.appendChild(preloadLink);
        }
    });
    
    function handleLoad() {
        isLoaded = true;
    }
    
    function handleError() {
        hasError = true;
    }
    
    // Determine final src (with fallback)
    $: finalSrc = hasError && fallback ? fallback : src;
</script>

<!-- Optimized img with intersection observer for lazy loading -->
<img
    bind:this={imgElement}
    src={finalSrc}
    {alt}
    {loading}
    class="{sizeMap[size]} object-contain select-none {classes}"
    class:opacity-0={!isLoaded && loading === 'lazy'}
    class:opacity-100={isLoaded || loading === 'eager'}
    style="transition: opacity 0.2s ease-in-out;"
    on:load={handleLoad}
    on:error={handleError}
    decoding="async"
    fetchpriority={preload ? 'high' : 'auto'}
/>