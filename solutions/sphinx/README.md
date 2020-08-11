### create a Sphinx project
#winpty docker run -it --rm -v /d/sphinx/docs:/docs sphinxdoc/sphinx sphinx-quickstart
docker run -it --rm -v D:\sphinx\docs:/docs sphinxdoc/sphinx sphinx-quickstart

### build HTML document
#winpty docker run --rm -v /d/sphinx/docs:/docs sphinxdoc/sphinx make html
docker run --rm -v D:\sphinx\docs:/docs sphinxdoc/sphinx make html