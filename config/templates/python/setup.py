# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


setup(
    name='{{.Name}}',
    version='0.0.0',
    author='{{.Author}}',
    author_email='{{.Email}}',

    license='MIT',

    zip_safe=True,
    packages=find_packages(exclude=['test', '.meta']),
)
