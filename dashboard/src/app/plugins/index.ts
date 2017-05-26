import auth from './auth';
import site from './site';

const plugins = [auth, site];

export default {
  components: plugins.reduce((a, p) => {
    return a.concat(p.routes.map((r) => r.component));
  }, []),
  routes: plugins.reduce((a, p) => {
    return a.concat(p.routes)
  }, [])
}
