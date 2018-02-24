# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


setup(
    name='pkgname',
    version='0.0.0',
    author='Marcus Veibäck',
    author_email='sirmar@gmail.com',

    description='{{description}}',
    url='https://github.com/sirmar/pkgname',
    license='MIT',

    zip_safe=True,
    packages=find_packages(exclude=['test', '.meta']),
)
