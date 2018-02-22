from unittest import TestCase
from nose.tools import istest
from pkgname.example import Example


class TestExample(TestCase):
    def setUp(self):
        self.example = Example()

    @istest
    def example_func_should_be_one(self):
        self.assertEqual(1, self.example.run())
