const manifest = (() => {
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
		client: {"start":"_app/immutable/entry/start.D4FhaIi5.js","app":"_app/immutable/entry/app.Dux7t5eW.js","imports":["_app/immutable/entry/start.D4FhaIi5.js","_app/immutable/chunks/entry.Cwgp6uEs.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.DGw2-XJO.js","_app/immutable/entry/app.Dux7t5eW.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./chunks/0-CWKxTb9C.js')),
			__memo(() => import('./chunks/1-BuwlOj97.js')),
			__memo(() => import('./chunks/2-DaZh6mM8.js')),
			__memo(() => import('./chunks/3-DUiWzezq.js')),
			__memo(() => import('./chunks/4-ZfN5jDXR.js')),
			__memo(() => import('./chunks/5-C80UDofR.js')),
			__memo(() => import('./chunks/6-CogDie9b.js')),
			__memo(() => import('./chunks/7-u_47S3k0.js')),
			__memo(() => import('./chunks/8-D6EoRtcX.js')),
			__memo(() => import('./chunks/9-CtGuRYC6.js'))
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
				id: "/home",
				pattern: /^\/home\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 9 },
				endpoint: null
			},
			{
				id: "/[event]/bracket",
				pattern: /^\/([^/]+?)\/bracket\/?$/,
				params: [{"name":"event","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,2,], errors: [1,,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/[event]/pools",
				pattern: /^\/([^/]+?)\/pools\/?$/,
				params: [{"name":"event","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,2,], errors: [1,,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/[event]/results",
				pattern: /^\/([^/]+?)\/results\/?$/,
				params: [{"name":"event","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,2,], errors: [1,,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/[event]/seeding",
				pattern: /^\/([^/]+?)\/seeding\/?$/,
				params: [{"name":"event","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,2,], errors: [1,,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/[event]/teams",
				pattern: /^\/([^/]+?)\/teams\/?$/,
				params: [{"name":"event","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,2,], errors: [1,,], leaf: 8 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();

const prerendered = new Set([]);

const base = "";

export { base, manifest, prerendered };
//# sourceMappingURL=manifest.js.map
