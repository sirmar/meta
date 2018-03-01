# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


setup(
    name='{{name}}',
    version='0.0.0',
    author='{{author}}',
    author_email='{{email}}',

    license='MIT',

    zip_safe=True,
    packages=find_packages(exclude=['test', '.meta']),
)
