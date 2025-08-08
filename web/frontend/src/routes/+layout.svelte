<script lang="ts">
    import Navbar from '$lib/components/shared/Navbar.svelte';
    import Footer from '$lib/components/shared/Footer.svelte';
    import '../global.css';
</script>

<div class="min-h-screen flex flex-col bg-theme-bg-primary"> 
    <Navbar />
    <main class="flex-grow"> 
        <slot />
    </main>
    <Footer />
</div>
