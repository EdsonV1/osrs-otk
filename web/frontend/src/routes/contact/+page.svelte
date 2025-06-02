<script lang="ts">

    const contactMethods = [
        {
            type: 'Email',
            value: 'edsonvillegas25@gmail.com',
            href: 'mailto:edsonvillegas25@gmail.com',
            icon: '✉️',
            description: 'For general inquiries, suggestions, or support.'
        },
        {
            type: 'GitHub',
            value: 'EdsonV1',
            href: 'https://github.com/EdsonV1',
            icon: '💻',
            description: 'Report issues, contribute, or check out the source code.'
        }
    ];
</script>

<div class="max-w-3xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
    <header class="mb-12 text-center">
        <h1 class="text-h1 text-theme-text-primary tracking-tight">
            Get In Touch
        </h1>
        <p class="mt-4 text-lg text-theme-text-secondary max-w-xl mx-auto">
            Have a question, suggestion, or just want to say hi? Here's how you can reach out.
        </p>
    </header>

    <div class="space-y-8">
        {#each contactMethods as method (method.type)}
            <div class="bg-theme-card-bg p-6 rounded-xl shadow-card border border-theme-border">
                <div class="flex items-start space-x-4">
                    {#if method.icon}
                        <div class="flex-shrink-0">
                            <span class="text-3xl">{method.icon}</span>
                        </div>
                    {/if}
                    <div class="flex-grow">
                        <h2 class="text-lg font-semibold text-theme-text-primary mb-1">
                            {method.type}
                        </h2>
                        {#if method.href !== '#'}
                            <a href={method.href}
                               target="_blank" rel="noopener noreferrer"
                               class="text-theme-accent hover:text-theme-accent-hover hover:underline break-all">
                                {method.value}
                            </a>
                        {:else}
                            <p class="text-theme-text-secondary break-all">{method.value}</p>
                        {/if}
                        {#if method.description}
                            <p class="mt-2 text-sm text-theme-text-tertiary">
                                {method.description}
                            </p>
                        {/if}
                    </div>
                </div>
            </div>
        {/each}
    </div>
</div>