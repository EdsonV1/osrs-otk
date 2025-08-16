<script lang="ts">
    import { page } from '$app/stores';
    
    let mobileMenuOpen = false;
    
    const navItems = [
        { href: '/', label: 'Home' },
        { href: '/skills', label: 'Skills' },
        { href: '/tools', label: 'Tools' },
        { href: '/contact', label: 'Contact' }
    ];
    
    function toggleMobileMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }
    
    function closeMobileMenu() {
        mobileMenuOpen = false;
    }
    
    function isActiveRoute(href: string): boolean {
        if (href === '/') {
            return $page.url.pathname === '/';
        }
        return $page.url.pathname.startsWith(href);
    }
</script>

<nav class="bg-theme-bg-secondary/95 backdrop-blur-lg border-b border-theme-border/50 shadow-lg sticky top-0 z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
            <!-- Logo -->
            <div class="flex items-center">
                <a href="/" class="flex items-center space-x-3 group" on:click={closeMobileMenu}>
                    <div class="w-8 h-8 bg-gradient-to-br from-theme-accent to-blue-500 rounded-lg flex items-center justify-center shadow-md group-hover:shadow-lg transition-all duration-200 group-hover:scale-105">
                        <span class="text-white font-bold text-sm">OS</span>
                    </div>
                    <span class="font-bold text-xl text-theme-text-primary group-hover:text-theme-accent transition-colors duration-200">
                        OSRS OTK
                    </span>
                </a>
            </div>
            
            <!-- Desktop Menu -->
            <div class="hidden md:flex items-center space-x-1">
                {#each navItems as item}
                    <a 
                        href={item.href} 
                        class="relative px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 group
                               {isActiveRoute(item.href) 
                                 ? 'text-theme-accent bg-theme-accent/10 shadow-sm' 
                                 : 'text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50'}"
                    >
                        {item.label}
                        {#if isActiveRoute(item.href)}
                            <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-1 h-1 bg-theme-accent rounded-full"></div>
                        {/if}
                        <div class="absolute inset-0 rounded-lg bg-theme-accent/5 opacity-0 group-hover:opacity-100 transition-opacity duration-200 -z-10"></div>
                    </a>
                {/each}
            </div>
            
            <!-- Mobile menu button -->
            <div class="md:hidden">
                <button 
                    type="button" 
                    class="relative p-2 rounded-lg text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50 focus:outline-none focus:ring-2 focus:ring-theme-accent/20 transition-all duration-200"
                    on:click={toggleMobileMenu}
                    aria-expanded={mobileMenuOpen}
                >
                    <span class="sr-only">Toggle main menu</span>
                    <div class="w-6 h-6 relative">
                        <span class="absolute block w-6 h-0.5 bg-current transform transition-all duration-300 {mobileMenuOpen ? 'rotate-45 top-3' : 'top-1'}"></span>
                        <span class="absolute block w-6 h-0.5 bg-current transform transition-all duration-300 {mobileMenuOpen ? 'opacity-0' : 'top-3'}"></span>
                        <span class="absolute block w-6 h-0.5 bg-current transform transition-all duration-300 {mobileMenuOpen ? '-rotate-45 top-3' : 'top-5'}"></span>
                    </div>
                </button>
            </div>
        </div>
        
        <!-- Mobile menu -->
        <div class="md:hidden {mobileMenuOpen ? 'block' : 'hidden'}">
            <div class="px-2 pt-2 pb-6 space-y-1 bg-theme-bg-secondary/95 backdrop-blur-lg border-t border-theme-border/30 mt-1">
                {#each navItems as item}
                    <a 
                        href={item.href}
                        class="block px-4 py-3 rounded-lg text-base font-medium transition-all duration-200
                               {isActiveRoute(item.href) 
                                 ? 'text-theme-accent bg-theme-accent/10 border-l-4 border-theme-accent shadow-sm' 
                                 : 'text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50'}"
                        on:click={closeMobileMenu}
                    >
                        {item.label}
                    </a>
                {/each}
            </div>
        </div>
    </div>
</nav>