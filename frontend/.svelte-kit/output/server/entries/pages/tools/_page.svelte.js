import { E as ensure_array_like, F as attr, J as stringify, I as escape_html } from "../../../chunks/index.js";
function _page($$payload) {
  const availableTools = [
    {
      name: "Ardougne Knight Calculator",
      description: "Optimize your Thieving training with Ardougne Knights: XP, GP, and efficiency.",
      href: "/tools/ardy-knights",
      iconSrc: "/images/tools/knight_of_ardougne.png"
    },
    {
      name: "Birdhouse Run Calculator",
      description: "Plan your birdhouse runs for maximum Hunter XP and valuable bird nests.",
      href: "/tools/birdhouses",
      iconSrc: "/images/birdhouse/redwood_bird_house.png"
    },
    {
      name: "Wintertodt Calculator",
      description: "Calculate Firemaking XP, Phoenix pet chances, and loot from Wintertodt rounds.",
      href: "/tools/wintertodt",
      iconSrc: "/images/skills/firemaking.png"
    },
    {
      name: "Guardians of the Rift Calculator",
      description: "Calculate Runecrafting XP rates, training time, rewards, and Abyssal Protector pet chances.",
      href: "/tools/gotr",
      iconSrc: "/images/skills/runecraft.png"
    }
  ];
  $$payload.out += `<div class="max-w-5xl mx-auto"><header class="mb-10 text-center"><h1 class="text-h1 text-theme-text-primary tracking-tight">OSRS Tools &amp; Calculators</h1> <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">A collection of utility tools and calculators to help with your Old School RuneScape journey.</p></header> `;
  if (availableTools.length > 0) {
    $$payload.out += "<!--[-->";
    const each_array = ensure_array_like(availableTools);
    $$payload.out += `<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6"><!--[-->`;
    for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
      let tool = each_array[$$index];
      $$payload.out += `<a${attr("href", tool.href)} class="group block p-6 bg-theme-card-bg rounded-xl shadow-card hover:shadow-inner-border transition-all duration-200 ease-in-out border border-theme-border hover:border-theme-accent focus:outline-none focus:ring-2 focus:ring-theme-accent focus:ring-offset-2 focus:ring-offset-theme-bg"><div class="flex items-center space-x-4 mb-3">`;
      if (tool.iconSrc) {
        $$payload.out += "<!--[-->";
        $$payload.out += `<img${attr("src", tool.iconSrc)}${attr("alt", `${stringify(tool.name)} icon`)} class="w-12 h-12 object-contain flex-shrink-0 mt-1 rounded-md group-hover:scale-105 transition-transform"/>`;
      } else {
        $$payload.out += "<!--[!-->";
      }
      $$payload.out += `<!--]--> <h2 class="text-lg font-semibold text-theme-text-primary group-hover:text-theme-accent transition-colors">${escape_html(tool.name)}</h2></div> <p class="text-sm text-theme-text-secondary leading-relaxed">${escape_html(tool.description)}</p></a>`;
    }
    $$payload.out += `<!--]--></div>`;
  } else {
    $$payload.out += "<!--[!-->";
    $$payload.out += `<div class="text-center py-10"><p class="text-theme-text-secondary text-lg">More tools coming soon!</p></div>`;
  }
  $$payload.out += `<!--]--></div>`;
}
export {
  _page as default
};
