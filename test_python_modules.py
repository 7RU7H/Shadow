
# Really want this to work intelligently

def test_modules():
    combined = list(sys.modules.items()) + [(dist.project_name, dist.version) for dist in pkg_resources.working_set]
    for item in combined:
        try:
            if len(item) == 2:
                module_name, module = item
                print(f"{module_name}: {module.__version__}")
            else:
                project_name, version = item[0], item[1]
                print(f"{project_name}: {version}")
        except AttributeError:
            print(f"{print_colors.WARNING}{module_name} does not have a __version__ attribute {print_colors.ENDC}")

