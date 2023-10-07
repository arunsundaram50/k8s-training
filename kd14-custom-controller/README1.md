Custom Resource Definitions (CRDs) in Kubernetes are used to extend the Kubernetes API by introducing new resource types. They're a powerful way to add custom functionalities and abstractions to a Kubernetes cluster. Here are situations in which one might consider writing a CRD:

1. **New Abstractions**: If you want to introduce a new abstraction that is not natively present in Kubernetes, a CRD can be helpful. For instance, the `Ingress` resource was originally a custom abstraction introduced via a CRD before it became a standard resource.

2. **Domain-Specific Configurations**: If your application or service has a unique configuration that doesn't map well to existing Kubernetes objects, you might use a CRD to represent that configuration.

3. **Operator Pattern**: The operator pattern is about extending Kubernetes to encode domain-specific knowledge into the cluster. An operator uses CRDs to represent high-level concepts and then ensures the desired state specified in these CRDs is maintained. For instance, a database operator might introduce a CRD for database clusters, allowing users to request databases easily.

4. **Managing External Resources**: If you want your Kubernetes cluster to manage resources outside of the cluster (like cloud resources: databases, storage, etc.), CRDs can be a way to represent these external resources. The controller for the CRD would then handle the provisioning and lifecycle of these external resources.

5. **Workflows & Pipelines**: For platforms that orchestrate complex workflows or pipelines (like CI/CD platforms), CRDs can be used to represent stages, tasks, or pipelines, providing a Kubernetes-native way to manage these workflows.

6. **Integration with Other Systems**: If you're looking to integrate Kubernetes with other systems (like monitoring, logging, or tracing platforms), CRDs can be used to define the integration points, rules, or configurations.

7. **Stateful Applications Management**: Managing stateful applications in Kubernetes can be complex. CRDs can be used to simplify this by defining higher-level abstractions that handle the nuances of stateful application management.

8. **Custom Validation & Behavior**: With CRDs, you can also specify custom validation rules for your resources. This allows you to ensure configurations meet specific criteria before being accepted by the API server.

9. **Custom Views & Groupings**: If you want to create custom views or groupings of existing resources, CRDs can be a way to represent these aggregated views.

10. **Consistency & Standardization**: If multiple teams or applications use common configurations or patterns, CRDs can ensure consistency and standardization across these configurations.

While CRDs are powerful, it's essential to assess whether you need a custom resource or if existing Kubernetes primitives (like ConfigMaps, Secrets, or custom annotations) can serve the purpose. Often, using native resources can reduce complexity. However, when the requirements are sufficiently unique or complex, CRDs provide a robust mechanism for extending Kubernetes to fit those needs.