<script lang="ts">
    import Navbar from '$lib/Navbar.svelte';
    import Footer from '$lib/Footer.svelte';
    import '../global.css';
</script>

<div class="min-h-screen flex flex-col bg-theme-bg"> 
    <Navbar />
    <div class="flex-grow container mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12"> 
        <slot />
    </div>
    <Footer />
</div>
