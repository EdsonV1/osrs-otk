import { E as ensure_array_like, F as attr, I as escape_html, G as attr_class, J as stringify } from "../../../chunks/index.js";
function _page($$payload) {
  const contactMethods = [
    {
      type: "Email",
      value: "edsonvillegas25@gmail.com",
      href: "mailto:edsonvillegas25@gmail.com",
      icon: "M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 4l-8 5-8-5V6l8 5 8-5v2z",
      description: "For general inquiries, suggestions, or support.",
      gradient: "from-blue-500 to-cyan-500"
    },
    {
      type: "GitHub",
      value: "EdsonV1",
      href: "https://github.com/EdsonV1",
      icon: "M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z",
      description: "Report issues, contribute, or check out the source code.",
      gradient: "from-gray-700 to-gray-900"
    }
  ];
  const stats = [
    {
      label: "Active Users",
      value: "2.5K+",
      icon: "M12 4.354a4 4 0 110 5.292V21a1 1 0 01-2 0v-11.354a4 4 0 110-5.292zM15 6a2 2 0 11-4 0 2 2 0 014 0zM6 16a2 2 0 11-4 0 2 2 0 014 0zM14 16a2 2 0 11-4 0 2 2 0 014 0z"
    },
    {
      label: "Calculators",
      value: "10+",
      icon: "M9 7h6l1.5-1.5L15 4H9l-1.5 1.5L9 7zm0 10l-1.5 1.5L9 20h6l1.5-1.5L15 17H9z"
    },
    {
      label: "Response Time",
      value: "< 24h",
      icon: "M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
    }
  ];
  const each_array = ensure_array_like(stats);
  const each_array_1 = ensure_array_like(contactMethods);
  $$payload.out += `<div class="relative bg-gradient-to-br from-theme-bg-primary via-theme-bg-secondary to-theme-bg-tertiary"><div class="absolute inset-0 bg-grid-pattern opacity-5"></div> <div class="relative max-w-7xl mx-auto py-16 px-4 sm:px-6 lg:px-8"><div class="text-center"><h1 class="text-4xl font-bold text-theme-text-primary sm:text-5xl lg:text-6xl"><span class="block">Let's</span> <span class="block bg-gradient-to-r from-blue-500 to-cyan-500 bg-clip-text text-transparent">Connect</span></h1> <p class="mt-6 text-xl text-theme-text-secondary max-w-3xl mx-auto leading-relaxed">Have questions about OSRS calculators? Need support? Want to contribute? <br class="hidden sm:block"/> I'm here to help and would love to hear from you.</p> <div class="mt-12 grid grid-cols-1 gap-4 sm:grid-cols-3"><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let stat = each_array[$$index];
    $$payload.out += `<div class="bg-theme-card-bg/50 backdrop-blur-sm p-6 rounded-xl border border-theme-border/50"><div class="flex items-center justify-center mb-2"><svg class="w-6 h-6 text-theme-accent" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"${attr("d", stat.icon)}></path></svg></div> <div class="text-2xl font-bold text-theme-text-primary">${escape_html(stat.value)}</div> <div class="text-sm text-theme-text-tertiary">${escape_html(stat.label)}</div></div>`;
  }
  $$payload.out += `<!--]--></div></div></div></div> <div class="max-w-6xl mx-auto py-16 px-4 sm:px-6 lg:px-8"><div class="text-center mb-12"><h2 class="text-3xl font-bold text-theme-text-primary mb-4">Get In Touch</h2> <p class="text-lg text-theme-text-secondary">Choose your preferred way to reach out</p></div> <div class="grid gap-8 md:grid-cols-2"><!--[-->`;
  for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
    let method = each_array_1[$$index_1];
    $$payload.out += `<div class="group relative"><div${attr_class(`absolute -inset-0.5 bg-gradient-to-r ${stringify(method.gradient)} rounded-xl blur opacity-30 group-hover:opacity-50 transition duration-300`)}></div> <div class="relative bg-theme-card-bg p-8 rounded-xl border border-theme-border group-hover:border-theme-accent/20 transition-all duration-300 hover:shadow-xl"><div class="flex items-start space-x-6"><div class="flex-shrink-0"><div${attr_class(`w-16 h-16 bg-gradient-to-r ${stringify(method.gradient)} rounded-xl flex items-center justify-center shadow-lg`)}><svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24"><path${attr("d", method.icon)}></path></svg></div></div> <div class="flex-grow min-w-0"><h3 class="text-xl font-semibold text-theme-text-primary mb-2 group-hover:text-theme-accent transition-colors duration-200">${escape_html(method.type)}</h3> <a${attr("href", method.href)} target="_blank" rel="noopener noreferrer" class="inline-flex items-center text-theme-text-secondary hover:text-theme-accent font-mono text-sm bg-theme-bg-secondary px-3 py-2 rounded-lg border border-theme-border hover:border-theme-accent/30 transition-all duration-200 group/link"><span class="truncate">${escape_html(method.value)}</span> <svg class="w-4 h-4 ml-2 opacity-0 group-hover/link:opacity-100 transform translate-x-0 group-hover/link:translate-x-1 transition-all duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path></svg></a> <p class="mt-3 text-theme-text-tertiary leading-relaxed">${escape_html(method.description)}</p></div></div></div></div>`;
  }
  $$payload.out += `<!--]--></div> <div class="mt-20 bg-theme-card-bg rounded-2xl p-8 border border-theme-border"><h3 class="text-2xl font-bold text-theme-text-primary mb-6 text-center">Frequently Asked Questions</h3> <div class="grid gap-6 md:grid-cols-2"><div class="space-y-4"><div><h4 class="font-semibold text-theme-text-primary mb-2">🚀 How quickly do you respond?</h4> <p class="text-sm text-theme-text-tertiary">Usually within 24 hours, often much faster for urgent issues.</p></div> <div><h4 class="font-semibold text-theme-text-primary mb-2">🔧 Can I contribute features?</h4> <p class="text-sm text-theme-text-tertiary">Absolutely! Check the GitHub repository for contribution guidelines.</p></div></div> <div class="space-y-4"><div><h4 class="font-semibold text-theme-text-primary mb-2">🐛 Found a bug?</h4> <p class="text-sm text-theme-text-tertiary">Please report it on GitHub with detailed steps to reproduce.</p></div> <div><h4 class="font-semibold text-theme-text-primary mb-2">💡 Have suggestions?</h4> <p class="text-sm text-theme-text-tertiary">I'd love to hear your ideas for new calculators or improvements!</p></div></div></div></div></div>`;
}
export {
  _page as default
};
