const API_URL = "https://awcu2nks23.us-east-1.awsapprunner.com";
const load = async ({ params }) => {
  const response = await fetch(`${API_URL}/api/${params.event}/results`);
  if (response.status == 404) {
    return {
      results: null
    };
  }
  const data = await response.json();
  return {
    results: data
  };
};

var _page_ts = /*#__PURE__*/Object.freeze({
  __proto__: null,
  load: load
});

const index = 6;
let component_cache;
const component = async () => component_cache ??= (await import('./_page.svelte-Xeabzk8o.js')).default;
const universal_id = "src/routes/[event]/results/+page.ts";
const imports = ["_app/immutable/nodes/6.CUNkk2Dj.js","_app/immutable/chunks/scheduler.Cdj48xKw.js","_app/immutable/chunks/index.krINOmAy.js","_app/immutable/chunks/TableHeadCell.DCsmDhEb.js","_app/immutable/chunks/bundle-mjs.C1WdQso_.js","_app/immutable/chunks/Heading.BAZVKAx1.js"];
const stylesheets = [];
const fonts = [];

export { component, fonts, imports, index, stylesheets, _page_ts as universal, universal_id };
//# sourceMappingURL=6-CogDie9b.js.map
