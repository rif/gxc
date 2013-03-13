from lxml import etree

schema = etree.XMLSchema(file='cgrates1.xsd')

parser = etree.XMLParser(schema = schema)

root = etree.parse('cgrates.xml', parser)
