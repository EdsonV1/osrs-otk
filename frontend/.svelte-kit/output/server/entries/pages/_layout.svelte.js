import { D as getContext, E as ensure_array_like, F as attr, G as attr_class, I as escape_html, J as stringify, C as pop, K as store_get, M as unsubscribe_stores, z as push, N as slot } from "../../chunks/index.js";
import "../../chunks/client.js";
import "clsx";
const getStores = () => {
  const stores$1 = getContext("__svelte__");
  return {
    /** @type {typeof page} */
    page: {
      subscribe: stores$1.page.subscribe
    },
    /** @type {typeof navigating} */
    navigating: {
      subscribe: stores$1.navigating.subscribe
    },
    /** @type {typeof updated} */
    updated: stores$1.updated
  };
};
const page = {
  subscribe(fn) {
    const store = getStores().page;
    return store.subscribe(fn);
  }
};
function Navbar($$payload, $$props) {
  push();
  var $$store_subs;
  let mobileMenuOpen = false;
  const navItems = [
    { href: "/", label: "Home" },
    { href: "/skills", label: "Skills" },
    { href: "/tools", label: "Tools" },
    { href: "/contact", label: "Contact" }
  ];
  function isActiveRoute(href) {
    if (href === "/") {
      return store_get($$store_subs ??= {}, "$page", page).url.pathname === "/";
    }
    return store_get($$store_subs ??= {}, "$page", page).url.pathname.startsWith(href);
  }
  const each_array = ensure_array_like(navItems);
  const each_array_1 = ensure_array_like(navItems);
  $$payload.out += `<nav class="bg-theme-bg-secondary/95 backdrop-blur-lg border-b border-theme-border/50 shadow-lg sticky top-0 z-50"><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"><div class="flex items-center justify-between h-16"><div class="flex items-center"><a href="/" class="flex items-center space-x-3 group"><div class="w-8 h-8 bg-gradient-to-br from-theme-accent to-blue-500 rounded-lg flex items-center justify-center shadow-md group-hover:shadow-lg transition-all duration-200 group-hover:scale-105"><span class="text-white font-bold text-sm">OS</span></div> <span class="font-bold text-xl text-theme-text-primary group-hover:text-theme-accent transition-colors duration-200">OSRS OTK</span></a></div> <div class="hidden md:flex items-center space-x-1"><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let item = each_array[$$index];
    $$payload.out += `<a${attr("href", item.href)}${attr_class(`relative px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 group ${stringify(isActiveRoute(item.href) ? "text-theme-accent bg-theme-accent/10 shadow-sm" : "text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50")}`)}>${escape_html(item.label)} `;
    if (isActiveRoute(item.href)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `<div class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-1 h-1 bg-theme-accent rounded-full"></div>`;
    } else {
      $$payload.out += "<!--[!-->";
    }
    $$payload.out += `<!--]--> <div class="absolute inset-0 rounded-lg bg-theme-accent/5 opacity-0 group-hover:opacity-100 transition-opacity duration-200 -z-10"></div></a>`;
  }
  $$payload.out += `<!--]--></div> <div class="md:hidden"><button type="button" class="relative p-2 rounded-lg text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50 focus:outline-none focus:ring-2 focus:ring-theme-accent/20 transition-all duration-200"${attr("aria-expanded", mobileMenuOpen)}><span class="sr-only">Toggle main menu</span> <div class="w-6 h-6 relative"><span${attr_class(`absolute block w-6 h-0.5 bg-current transform transition-all duration-300 ${stringify("top-1")}`)}></span> <span${attr_class(`absolute block w-6 h-0.5 bg-current transform transition-all duration-300 ${stringify("top-3")}`)}></span> <span${attr_class(`absolute block w-6 h-0.5 bg-current transform transition-all duration-300 ${stringify("top-5")}`)}></span></div></button></div></div> <div${attr_class(`md:hidden ${stringify("hidden")}`)}><div class="px-2 pt-2 pb-6 space-y-1 bg-theme-bg-secondary/95 backdrop-blur-lg border-t border-theme-border/30 mt-1"><!--[-->`;
  for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
    let item = each_array_1[$$index_1];
    $$payload.out += `<a${attr("href", item.href)}${attr_class(`block px-4 py-3 rounded-lg text-base font-medium transition-all duration-200 ${stringify(isActiveRoute(item.href) ? "text-theme-accent bg-theme-accent/10 border-l-4 border-theme-accent shadow-sm" : "text-theme-text-secondary hover:text-theme-text-primary hover:bg-theme-bg-primary/50")}`)}>${escape_html(item.label)}</a>`;
  }
  $$payload.out += `<!--]--></div></div></div></nav>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
function Footer($$payload, $$props) {
  push();
  $$payload.out += `<footer class="bg-theme-bg-secondary border-t border-theme-border mt-auto"><div class="max-w-7xl mx-auto py-5 px-4 sm:px-6 lg:px-8 text-sm text-theme-text-secondary"><div class="flex flex-col md:flex-row justify-between items-center md:items-start space-y-3 md:space-y-0"><div class="text-center md:text-left"><p>© ${escape_html((/* @__PURE__ */ new Date()).getFullYear())} OSRS OTK. Developed by <a href="mailto:edsonvillegas25@gmail.com" target="_blank" rel="noopener noreferrer" class="font-semibold text-theme-text-primary hover:text-theme-accent transition-colors">Edson</a>.</p> <p class="text-xs text-theme-text-tertiary mt-0.5">Data sourced from various community APIs &amp; resources.</p></div> <div class="flex items-center space-x-5 shrink-0 mt-3 md:mt-0"><a href="/contact" class="flex items-center hover:text-theme-accent transition-colors duration-150"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-envelope" viewBox="0 0 16 16"><path d="M0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2zm2-1a1 1 0 0 0-1 1v.217l7 4.2 7-4.2V4a1 1 0 0 0-1-1zm13 2.383-4.708 2.825L15 11.105zm-.034 6.876-5.64-3.471L8 9.583l-1.326-.795-5.64 3.47A1 1 0 0 0 2 13h12a1 1 0 0 0 .966-.741M1 11.105l4.708-2.897L1 5.383z"></path></svg> <span class="ml-1.5">Contact</span></a> `;
  {
    $$payload.out += "<!--[-->";
    $$payload.out += `<a href="https://github.com/EdsonV1" target="_blank" rel="noopener noreferrer" class="flex items-center hover:text-theme-accent transition-colors duration-150"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0 0 16 8c0-4.42-3.58-8-8-8Z"></path></svg> <span class="ml-1.5">GitHub</span></a>`;
  }
  $$payload.out += `<!--]--></div></div> <div class="mt-4 pt-4 border-t border-theme-border border-opacity-40 text-center text-xs text-theme-text-tertiary"><p>This is a fan-made project and is not affiliated with Jagex Ltd. Old School RuneScape™ is a trademark of Jagex Ltd.</p></div></div></footer>`;
  pop();
}
function _layout($$payload, $$props) {
  $$payload.out += `<div class="min-h-screen flex flex-col bg-theme-bg-primary">`;
  Navbar($$payload);
  $$payload.out += `<!----> <main class="flex-grow"><!---->`;
  slot($$payload, $$props, "default", {});
  $$payload.out += `<!----></main> `;
  Footer($$payload);
  $$payload.out += `<!----></div>`;
}
export {
  _layout as default
};
