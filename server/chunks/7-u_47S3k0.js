const API_URL$1 = "https://awcu2nks23.us-east-1.awsapprunner.com";
const load = async ({ params }) => {
  const response = await fetch(`${API_URL$1}/api/${params.event}/seeding`);
  if (response.status == 404) {
    return {
      seeding: null
    };
  }
  const data = await response.json();
  return {
    seeding: data
  };
};

var _page_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const API_URL = "https://awcu2nks23.us-east-1.awsapprunner.com";
const actions = {
  start: async ({ params }) => {
    await fetch(`${API_URL}/api/${params.event}/start-elimination`, {
      method: "POST"
    });
  }
};

var _page_server_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  actions: actions
});

const index = 7;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-oRQk2giQ.js')).default;
const universal_id = "src/routes/[event]/seeding/+page.ts";
const server_id = "src/routes/[event]/seeding/+page.server.ts";
const imports = ["_app/immutable/nodes/7.DpW_Iz5i.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/TableHeadCell.DCsmDhEb.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/Heading.BAZVKAx1.js","_app/immutable/chunks/admin.D1QBpGg2.js","_app/immutable/chunks/index.DGw2-XJO.js","_app/immutable/chunks/forms.CYrJzw3n.js","_app/immutable/chunks/entry.Cwgp6uEs.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, _page_server_ts as server, server_id, stylesheets, _page_ts as universal, universal_id };
//# sourceMappingURL=7-u_47S3k0.js.map
