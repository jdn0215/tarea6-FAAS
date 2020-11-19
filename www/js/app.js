const URL = "https://ecstatic-lumiere-612ea1.netlify.app/api/";

function getAll(entity) {
    fetch(URL + entity)
        .then((response) => response.json())
        .then((data) => {
            fetch('/template/list/' + entity + '.html')
                .then((response) => response.text())
                .then((template) => {
                    var rendered = Mustache.render(template, data);
                    document.getElementById('content').innerHTML = rendered;
                });
        })
}

function getById(query, entity) {
    var params = new URLSearchParams(query);
    fetch(URL + entity + '/?id=' + params.get('id'))
        .then((response) => response.json())
        .then((data) => {
            fetch('/template/detail/' + entity + '.html')
                .then((response) => response.text())
                .then((template) => {
                    var rendered = Mustache.render(template, data);
                    document.getElementById('content').innerHTML = rendered;
                });
        })
}

function home() {
    fetch('/template/home.html')
        .then((response) => response.text())
        .then((template) => {
            var rendered = Mustache.render(template, {});
            document.getElementById('content').innerHTML = rendered;
        });
}

function init() {
    router = new Navigo(null, false, '#!');
    router.on({
        '/facturas': function() {
            getAll('facturas');
        },
        '/clientes': function() {
            getAll('clientes');
        },
        '/productos': function() {
            getAll('productos');
        },
        '/facturaById': function(_, query) {
            getById(query, 'facturas');
        },
        '/clienteById': function(_, query) {
            getById(query, 'clientes');
        },
        '/productoById': function(_, query) {
            getById(query, 'productos');
        }
    });
    router.on(() => home());
    router.resolve();
}