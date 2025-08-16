export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {start:"_app/immutable/entry/start.C-mZTvHA.js",app:"_app/immutable/entry/app.DbueP4HA.js",imports:["_app/immutable/entry/start.C-mZTvHA.js","_app/immutable/chunks/0OQ3iKNu.js","_app/immutable/chunks/CZMTbaSA.js","_app/immutable/chunks/BQSnFksB.js","_app/immutable/chunks/CYgJF_JY.js","_app/immutable/entry/app.DbueP4HA.js","_app/immutable/chunks/CZMTbaSA.js","_app/immutable/chunks/DtbfWuKm.js","_app/immutable/chunks/CFX-n10L.js","_app/immutable/chunks/BQSnFksB.js","_app/immutable/chunks/Bea9AwOS.js","_app/immutable/chunks/BZk283XO.js","_app/immutable/chunks/DtU3kO7g.js","_app/immutable/chunks/DXSNkYra.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js')),
			__memo(() => import('./nodes/5.js')),
			__memo(() => import('./nodes/6.js')),
			__memo(() => import('./nodes/7.js')),
			__memo(() => import('./nodes/8.js')),
			__memo(() => import('./nodes/9.js')),
			__memo(() => import('./nodes/10.js')),
			__memo(() => import('./nodes/11.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/contact",
				pattern: /^\/contact\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/skills",
				pattern: /^\/skills\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/skills/[skill_name]",
				pattern: /^\/skills\/([^/]+?)\/?$/,
				params: [{"name":"skill_name","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/tools",
				pattern: /^\/tools\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/tools/ardy-knights",
				pattern: /^\/tools\/ardy-knights\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 8 },
				endpoint: null
			},
			{
				id: "/tools/birdhouses",
				pattern: /^\/tools\/birdhouses\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 9 },
				endpoint: null
			},
			{
				id: "/tools/gotr",
				pattern: /^\/tools\/gotr\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 10 },
				endpoint: null
			},
			{
				id: "/tools/wintertodt",
				pattern: /^\/tools\/wintertodt\/?$/,
				params: [],
				page: { layouts: [0,2,], errors: [1,,], leaf: 11 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
