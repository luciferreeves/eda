from configparser import ConfigParser
from django.template import Origin
from django.template.loaders.filesystem import Loader as FileSystemLoader
from django.contrib.staticfiles.finders import FileSystemFinder
from django.conf import settings
import os


class ThemeLoader(FileSystemLoader):
    def get_template_sources(self, template_name, template_dirs=None):
        theme = getattr(settings, "ACTIVE_THEME", "default")
        theme_dir = os.path.join(settings.BASE_DIR, "themes", theme, "templates")
        full_path = os.path.join(theme_dir, template_name)
        yield Origin(
            name=full_path,
            template_name=template_name,
            loader=self,
        )


class ThemeStaticFinder(FileSystemFinder):
    def __init__(self, app_names=None, *args, **kwargs):
        self.locations = []
        self.storages = {}
        theme = getattr(settings, "ACTIVE_THEME", "default")
        theme_static_dir = os.path.join(settings.BASE_DIR, "themes", theme, "static")
        self.locations.append(("", theme_static_dir))
        super(FileSystemFinder, self).__init__(*args, **kwargs)


def get_theme_config():
    theme = getattr(settings, "ACTIVE_THEME", "default")
    theme_dir = os.path.join(settings.BASE_DIR, "themes", theme)
    config_path = os.path.join(theme_dir, "theme.ini")

    config = ConfigParser()
    config.read(config_path)

    return {
        "name": config.get("theme", "name", fallback="Unknown Theme"),
        "description": config.get(
            "theme", "description", fallback="No description provided"
        ),
        "version": config.get("theme", "version", fallback="0.0.1"),
        "author": config.get("theme", "author", fallback="Unknown Author"),
    }
