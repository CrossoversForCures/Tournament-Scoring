const API_URL = "https://awcu2nks23.us-east-1.awsapprunner.com";
const load = async ({ params }) => {
  const response = await fetch(`${API_URL}/api/${params.event}/teams`);
  if (response.status == 404) {
    return {
      teams: null
    };
  }
  const data = await response.json();
  return {
    teams: data
  };
};

var _page_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 8;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-JNq1RXZI.js')).default;
const universal_id = "src/routes/[event]/teams/+page.ts";
const imports = ["_app/immutable/nodes/8.Dn_e81gK.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/TableHeadCell.DCsmDhEb.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/Heading.BAZVKAx1.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, stylesheets, _page_ts as universal, universal_id };
//# sourceMappingURL=8-D6EoRtcX.js.map
