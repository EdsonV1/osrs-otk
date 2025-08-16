import * as universal from '../entries/pages/skills/_skill_name_/_page.ts.js';

export const index = 6;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/skills/_skill_name_/_page.svelte.js')).default;
export { universal };
export const universal_id = "src/routes/skills/[skill_name]/+page.ts";
export const imports = ["_app/immutable/nodes/6.C_khRwpp.js","_app/immutable/chunks/CYgJF_JY.js","_app/immutable/chunks/CFX-n10L.js","_app/immutable/chunks/CZMTbaSA.js","_app/immutable/chunks/CCMPKMQG.js","_app/immutable/chunks/DtbfWuKm.js","_app/immutable/chunks/Bea9AwOS.js","_app/immutable/chunks/DY3kE1xq.js","_app/immutable/chunks/Cr0IC5xA.js","_app/immutable/chunks/4U71Av4S.js","_app/immutable/chunks/BMgiCmxP.js","_app/immutable/chunks/DtU3kO7g.js","_app/immutable/chunks/DXSNkYra.js","_app/immutable/chunks/BQSnFksB.js","_app/immutable/chunks/CjRuUf64.js"];
export const stylesheets = [];
export const fonts = [];
